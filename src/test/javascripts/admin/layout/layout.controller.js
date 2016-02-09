(function () {
    'use strict';

    angular
        .module('app.layout')
        .controller('Layout', Layout);

    Layout.$inject = [
    	'logger', '$rootScope', '$state', '$filter', 'config'
    ];

    function Layout(logger, $rootScope, $state, $filter, config) {
        /*jshint validthis: true */
        var vm = this;
        // Variables

        // Methods

        // Activate
		activate();
        ///////////

        /**
         * Activate Method
         */
        function activate() {
            getNavRoutes();
            registerBroadcast();
            $rootScope.$broadcast('auth:updated');
            $rootScope.srefToHref = srefToHref;
            $rootScope.config = config;
        }

        /**
         * Get and check if has menu.
         */
        function getNavRoutes() {
            var menuTop  = [];
            var menuSide = [];
            $state.get().forEach(function (item) {
                if (item.settings && item.settings.text) {
                    var settings = item.settings || {show: true};
                    var permissions = settings.permissions || {};
                    //if ($auth.canAccess(permissions)) {
                    //    if (!!settings.isub) {
                    //        subItems.push(settings);
                    //    } else {
                    //        menuItems.push(settings);
                    //    }
                    //}

                    if (settings && settings.show_side === true) {
                        menuSide.push(settings);
                    }
                    if (settings && settings.show_top === true) {
                        menuTop.push(settings);
                    }
                }
            });

            $rootScope.menuSide = $filter('orderBy')(menuSide, '+order', false);
            $rootScope.menuTop  = $filter('orderBy')(menuTop, '+order', false);
        }

         /**
         * Register events broadcast
         */
        function registerBroadcast() {
            $rootScope.$on('view:loading', function (nil, action){
                //if (['start', 'complete'].indexOf(action) != -1) {
                //    cfpLoadingBar[action]();
                //} else {
                //    logger.error('Action for loading bar is not defined!');
                //}
                logger.debug(action);
            });
        }

        /**
         * Return thr url form alias router
         * @param   sref
         * @returns string
         */
        function srefToHref(sref) {
            return $state.href(sref);
        }
    }
})();
