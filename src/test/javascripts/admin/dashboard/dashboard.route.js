(function () {
    'use strict';

    angular
        .module('app.dashboard')
        .run(routeConfig);

    routeConfig.$inject = ['routehelper', 'routehelperConfig'];
    /* @ngInject */
    function routeConfig(routehelper) {
        routehelper.configureRoutes(getRoutes());
    }

    function getRoutes() {
        return [
            {
                state: 'grid.dashboard',
                config: {
                    url: '/dashboard',

                    templateUrl: 'admin/dashboard/dashboard.html',
                    title: 'Dashboard',

                    controller: 'Dashboard',
                    controllerAs: 'vm',

                    resolve: {},

                    settings: {
                        show_side: true,
                        permissions: {
                            //except: ['anonymous'],
                            //redirectTo: 'page.login'
                        },

                        order: 1,
                        text: 'Dashboard',
                        sref: 'grid.dashboard',
                        icon: 'fa fa-dashboard',
                        translate: 'menu.dashboard'
                    }
                }
            }
        ];
    }

})();
