(function () {
    'use strict';

    angular
        .module('app.fourOhFour')
        .run(routeConfig);

    routeConfig.$inject = ['routehelper', 'routehelperConfig'];

    /* @ngInject */
    function routeConfig(routehelper, routehelperConfig) {
        routehelper.configureRoutes(getRoutes(routehelperConfig));
    }

    function getRoutes(helper) {
        return [
            {
                state: 'empty.fourOhFour',
                config: {
                    url: '/fourOhFour',

                    templateUrl: 'admin/fourOhFour/fourOhFour.html',
                    title: 'FourOhFour',

                    controller: 'FourOhFour',
                    controllerAs: 'vm',

                    resolve: {},

                    settings: {
                        permissions: {
                        }
                    }
                }
            }
        ];
    }

})();
