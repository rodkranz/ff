(function () {
    'use strict';

    angular
        .module('app.school')
        .run(routeConfig);

    routeConfig.$inject = ['routehelper'];
    /* @ngInject */
    function routeConfig(routehelper) {
        routehelper.configureRoutes(getRoutes());
    }

    function getRoutes() {
        return [
            {
                state: 'grid.school',
                config: {
                    url: '/school',

                    templateUrl: 'admin/school/schoolList.html',
                    title: 'School List',

                    controller: 'SchoolList',
                    controllerAs: 'vm',

                    settings: {
                        show_side: true,
                        permissions: {
                        },

                        order: 1,
                        text: 'School',
                        sref: 'grid.school',
                        icon: 'fa fa-building',
                        translate: 'menu.school'
                    }
                }
            }, {
                state: 'grid.schoolNew',
                config: {
                    url: '/school/new',

                    templateUrl: 'admin/school/schoolForm.html',
                    title: 'New School',

                    controller: 'SchoolForm',
                    controllerAs: 'vm',

                    resolve: {
                        school: function () {
                            return {}
                        }
                    },

                    settings: {
                        permissions: {
                        }
                    }
                }
            }, {
                state: 'grid.schoolEdit',
                config: {
                    url: '/school/:id/edit',

                    templateUrl: 'admin/school/schoolForm.html',
                    title: 'School Editing',

                    controller: 'SchoolForm',
                    controllerAs: 'vm',

                    resolve: {
                        school: function(datacontext, $stateParams) {
                            return datacontext
                                .school
                                .findSchoolById($stateParams.id);
                        }
                    },

                    settings: {
                        permissions: {
                        }
                    }
                }
            }
        ];
    }

})();
