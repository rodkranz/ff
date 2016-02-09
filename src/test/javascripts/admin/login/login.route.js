(function () {
    'use strict';

    angular
        .module('app.login')
        .run(routeConfig);

    routeConfig.$inject = ['routehelper', 'routehelperConfig'];
    /* @ngInject */
    function routeConfig(routehelper, routehelperConfig) {
        routehelper.configureRoutes(getRoutes(routehelperConfig));
    }

    function getRoutes(helper) {
        return [
            {
                state: 'empty.login',
                config: {
                    url: '/login',

                    templateUrl: 'admin/login/login.html',
                    title: 'Login',

                    controller: 'Login',
                    controllerAs: 'vm',

                    resolve: {},

                    settings: {
                        permissions: {},
                        order: 1,
                        text: 'Login',
                        sref: 'app.login',
                        icon: 'fa fa-dashboard',
                        translate: 'menu.login'
                    }
                }
            }
        ];
    }

})();
