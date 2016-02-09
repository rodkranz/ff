(function () {
    'use strict';

    var core = angular.module('app.core');

    var config = {
        debug: true,

        version: '1.1.0',
        prefix: '[EzWayApp]',
        timezone: 'Europe/lisbon',

        appName: 'EzWayApp',
        slogan: 'EzWayApp AngularJs',
        apiHost: 'http://localhost:3000/api/',
        // apiHost: 'http://www.ezwayapp.com/api/',
        apiVersion: 'v1/',

        i18nFolder: '/assets/i18n/',

        timeout: 2500,
        toastPosition: 'toast-bottom-right',

        pathDefault: '/web/login',
        pathNotFound: '/web/fourOhFour',
        pathDenied: '/web/denied',

        dateFormat: 'DD/MM/YYYY',

        whiteList: [
            "self",
            "http*://*localhost*",
            "http*://*ezwayapp*"
        ],

        repositories: [
            'user', 'dailyPost', 'school', 'schoolGalleries', 'schoolComments'
        ],

        language: {
            default: "eng",
            available: {
                eng: "English"
            }
        }
    };

    // set config in constants
    core.constant('config', config);

    // configurations
    core.config(configure);

    configure.$inject = [
        '$logProvider', '$stateProvider', '$urlRouterProvider',
        'routehelperConfigProvider', 'exceptionHandlerProvider',
        '$locationProvider', 'angularMomentConfig', 'toastr',
        '$translateProvider', '$httpProvider', '$compileProvider',
        'RestangularProvider', '$sceDelegateProvider'
    ];

    /* @ngInject */
    function configure($logProvider, $stateProvider, $urlRouterProvider,
                       routehelperConfigProvider, exceptionHandlerProvider,
                       $locationProvider, angularMomentConfig, toastr,
                       $translateProvider, $httpProvider, $compileProvider,
                       RestangularProvider, $sceDelegateProvider) {
        /* jshint validthis:true */

        // Initial configuration
        activate();
        ///////////

        //
        function activate() {
            configureLogging(); // configuration of log
            configureException();
            configureRouter();
            configureLocation();
            configureTranslate();
            //configureCompile();
            configureRestangular();
            configureMoment();
            configureToastr();
            configureWhiteList();
        }

        //configure debug mode on compile.
        function configureCompile() {
            $compileProvider.debugInfoEnabled(config.debug);
        }

        //log configuration
        function configureLogging() {
            if ($logProvider.debugEnabled) {
                $logProvider.debugEnabled(config.debug);
            }
        }

        //configure prefix
        function configureLocation() {
            //$locationProvider.html5Mode(true).hashPrefix('!');
            $locationProvider.hashPrefix('!');
        }

        //configure Restangular
        function configureRestangular() {
            RestangularProvider.setBaseUrl(config.apiHost + config.apiVersion);
            RestangularProvider.addResponseInterceptor(
                function (data, operation, what, url, response, deferred) {
                    return data[response.headers('x-resource')];
                }
            );
        }

        //configure translate
        function configureTranslate() {
            $translateProvider.useStaticFilesLoader({
                prefix: '/application/angular_language?q=',
                suffix: '.json'
            });

            $translateProvider.preferredLanguage(config.language.default);
            $translateProvider.useLocalStorage();
            $translateProvider.usePostCompiling(true);
            $translateProvider.useSanitizeValueStrategy('sanitize');
        }

        //configure moment
        function configureMoment() {
            angularMomentConfig = {
                dateFormat: config.dateFormat,
                preprocess: 'unix',
                timezone: config.timezone || 'Europe/London'
            };
        }

        //configure toastr
        function configureToastr() {
            toastr.options.timeOut = config.timeout;
            toastr.options.positionClass = config.toastPosition;
        }

        // set router configure.
        function configureRouter() {
            // Configure the common route provider
            routehelperConfigProvider.config.$stateProvider = $stateProvider;
            routehelperConfigProvider.config.$urlRouterProvider = $urlRouterProvider;
            routehelperConfigProvider.config.$stateProvider = $stateProvider;
            routehelperConfigProvider.config.docTitle = config.prefix + ' ';
            var resolveAlways = {
                ready: function (datacontext) {
                    return datacontext.ready();
                }
            };
            routehelperConfigProvider.config.resolveAlways = resolveAlways;
        }

        // exception configure.
        function configureException() {
            // Configure the common exception handler
            exceptionHandlerProvider.configure(config.appErrorPrefix);
        }

        // write white list
        function configureWhiteList() {
            $sceDelegateProvider.resourceUrlWhitelist = config.whiteList;
        }
    }
})();
