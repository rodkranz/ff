(function () {
    'use strict';

    angular
        .module('app.layout')
        .run(appRun);

    appRun.$inject = ['routehelper'];

    /* @ngInject */
    function appRun(routehelper) {
        routehelper.configureRoutes(getRoutes());
    }

    function getRoutes() {
        return [
            {
                state: 'grid',
                config: {
                    url: '/app',
                    abstract: true,
                    controller: 'Grid',
                    templateUrl: 'admin/layout/grid/grid.html',
                    data: {
                        permissions: {
                            private: true
                        }
                    }
                }
            },
            {
                state: 'empty',
                config: {
                    url: '/web',
                    abstract: true,
                    controller: 'Empty',
                    templateUrl: 'admin/layout/empty/empty.html'
                }
            }
        ];
    }
})();
