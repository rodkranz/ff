(function () {
    'use strict';

    angular
        .module('app.data')
        .factory('repository.dailyPost', dailyPost);

    dailyPost.$inject = ['$rest'];

    /*@ngInject*/
    function dailyPost($rest) {
        /* jshint validthis:true */

        var resourceName        = "daily_posts";
        var resourceServices = {
            findAllDailyPost:   findAllDailyPost,
            findDailyPostById:  findDailyPostById,
            saveDailyPost:      saveDailyPost,
            removeDailyPost:    removeDailyPost
        };

        return resourceServices;

        function findAllDailyPost(params) {
            return $rest
                .restFulAuth
                .all(resourceName)
                .getList(params);
        }

        function findDailyPostById(id) {
            return $rest
                .restAuth
                .one(resourceName, id)
                .get();
        }

        function saveDailyPost(dailyPost) {
            if ( dailyPost.restangularized ) {
                return dailyPost.save();
            } else {
                return $rest
                    .restFulAuth
                    .all(resourceName)
                    .post(dailyPost);
            }
        }

        function removeDailyPost(id) {
            return $rest
                .restFulAuth
                .one(resourceName, id)
                .remove();
        }
    }

})();