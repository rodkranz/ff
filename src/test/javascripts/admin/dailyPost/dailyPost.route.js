(function () {
    'use strict';

    angular
        .module('app.dailyPost')
        .run(routeConfig);

    routeConfig.$inject = ['routehelper', 'routehelperConfig'];
    /* @ngInject */
    function routeConfig(routehelper) {
        routehelper.configureRoutes(getRoutes());
    }

    function getRoutes() {
        return [
            {
                state: 'grid.dailyPostList',
                config: {
                    url: '/dailyPost',

                    templateUrl: 'admin/dailyPost/dailyPostList.html',
                    title: 'Daily Post List',

                    controller: 'DailyPostList',
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
                        sref: 'grid.dailyPostList',
                        icon: 'fa fa-pencil-square-o',
                        translate: 'menu.daily_post'
                    }
                }
            },
            {
                state: 'grid.dailyPostNew',
                config: {
                    url: '/dailyPost/create',

                    templateUrl: 'admin/dailyPost/dailyPostForm.html',
                    title: 'Daily Post List',

                    controller: 'DailyPostForm',
                    controllerAs: 'vm',

                    resolve: {
                        dailyPost: function(){
                            return {};
                        }
                    }
                }
            },
            {
                state: 'grid.dailyPostEdit',
                config: {
                    url: '/dailyPost/:id/edit',

                    templateUrl: 'admin/dailyPost/dailyPostForm.html',
                    title: 'Daily Post List',

                    controller: 'DailyPostForm',
                    controllerAs: 'vm',

                    resolve: {
                        dailyPost: function(datacontext, $stateParams){
                            return datacontext
                                .dailyPost
                                .findDailyPostById($stateParams.id);
                        }
                    }
                }
            }
        ];
    }

})();
