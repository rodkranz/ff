(function() {
    'use strict';

    angular
        .module('blocks.router')
        .provider('routehelperConfig', routehelperConfig)
        .factory('routehelper', routehelper);

    // Must configure via the routehelperConfigProvider
    function routehelperConfig() {
        /* jshint validthis:true */
        this.config = {
            // These are the properties we need to set
            // $routeProvider: undefined
            // docTitle: ''
            // resolveAlways: {ready: function(){ } }
        };

        this.$get = function() {
            return {
                config: this.config,
                item:   {}
            };
        };
    }

    routehelper.$inject = [
        '$location', '$rootScope', '$window', '$timeout', 'logger',
        'routehelperConfig', '$state', '$auth', 'config'
    ];

    function routehelper($location, $rootScope, $window, $timeout, logger,
                         routehelperConfig, $state, $auth, config) {

        var handlingRouteChangeError = false;
        var routeCounts = {
            notFound: 0,
            error: 0,
            change: 0
        };

        var routes = [];
        //var $routeProvider = routehelperConfig.config.$routeProvider;
        var $stateProvider     = routehelperConfig.config.$stateProvider;
        var $urlRouterProvider = routehelperConfig.config.$urlRouterProvider;

        var service = {
            configureRoutes: configureRoutes,
            getRoutes: getRoutes,
            routeCounts: routeCounts
        };

        init();

        return service;
        ///////////////

        function init() {
            handleRoutingStart();
            handleRoutingErrors();
            handleRoutingNotFound();
            updateDocTitle();
        }


        function configureRoutes(routes) {

            routes.forEach(function (route) {
                route.config.resolve =
                    angular.extend(route.config.resolve || {
                            $item: function () {
                                return null;
                            }
                        },
                        routehelperConfig.config.resolveAlways);
                $stateProvider.state(route.state, route.config);
            });

            $urlRouterProvider.otherwise( config.pathDefault );
        }

        function handleRoutingNotFound() {
            // Hook not found
            $rootScope.$on('$stateNotFound', function (event, unfoundState, fromState, fromParams) {
                logger.error('Page not found!');
                routeCounts.notFound++;
                $location.path(config.pathNotFound);
            });
        }

        function handleRoutingErrors() {
            // Hook error
            $rootScope.$on('$stateChangeError', function (event, toState, toParams, fromState, fromParams, error) {

                if( error ) {
                    logger.error(error);
                    return;
                }

                routeCounts.errors++;
                handlingRouteChangeError = true;


                var destination = (current && (current.title || current.name || current.loadedTemplateUrl)) ||
                    'unknown target';

                var msg = 'Error routing to ' + destination + '. ' + (rejection.msg || '');
                logger.warning(msg, [current]);
                $location.path(config.pathNotFound);
            });
        }

        function handleRoutingStart() {
            $rootScope.$on('$stateChangeStart', function (event, toState, toParams, fromState, fromParams) {
                $rootScope.$emit('view:loading', 'start');
                ga('send', 'pageview', toState.url);
            });
        }

        function getRoutes() {
            for (var prop in $route.routes) {
                if ($route.routes.hasOwnProperty(prop)) {
                    var route = $route.routes[prop];
                    var isRoute = !!route.title;
                    if (isRoute) {
                        routes.push(route);
                    }
                }
            }
            return routes;
        }

        function updateDocTitle() {
            $rootScope.$on('$stateChangeSuccess', function (event, toState, toParams, fromState, fromParams) {
                $window.scrollTo(0, 0);

                routeCounts.changes++;
                handlingRouteChangeError = false;
                var title = routehelperConfig.config.docTitle + ' ' + (toState && toState.title || '');
                $rootScope.title = title; // data bind to <title>

                event.targetScope.$watch('$viewContentLoaded', function () {
                    $rootScope.$emit('view:loading', 'complete');
                });
            });
        }
    }
})();
