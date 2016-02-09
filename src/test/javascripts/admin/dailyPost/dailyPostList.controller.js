(function () {
    'use strict';

    angular
        .module('app.dailyPost')
        .controller('DailyPostList', DailyPostList);

    DailyPostList.$inject = [
        'logger', 'datacontext', 'NgTableParams', '$filter', 'swal'
    ];

    /* @ngInject */
    function DailyPostList(logger, datacontext, NgTableParams, $filter, swal) {
        /* jshint validthis:true */
        var vm = this;

        // Variable

        // Methods
        vm.removeDailyPost = removeDailyPost;

        activate();
        ///////////

        /* Methods */
        function activate(){
            logger.debug('Admin > Daily Post Form > Loaded');
            makeTable();
        }


        function makeTable () {
            vm.tableParams = new NgTableParams({
                page: 1,
                count: 10
            }, {
                total: 0,
                getData: function(params) {
                    var parameters = params.url() || {};
                    return datacontext
                        .dailyPost
                        .findAllDailyPost(parameters)
                        .then(setData);

                    function setData(response) {
                        params.total(response.headers('x-total'));
                        return response.data;
                    }
                }
            });
        }

        function removeDailyPost(dailyPost) {
            swal({
                title: $filter('translate')('alert.are_you_sure'),
                text:  $filter('translate')('alert.are_you_sure_that_want_to_delete_this_item'),
                type:  "warning",
                showCancelButton: true,
                confirmButtonColor: "#DD6B55",
                confirmButtonText: $filter('translate')('button.yes')
            }, confirmDelete);

            function confirmDelete (){
                datacontext
                    .dailyPost
                    .removeDailyPost(dailyPost.id)
                    .then(removeSuccess, removeFail);

                function removeSuccess() {
                    logger.success('alert.daily_post_has_been_deleted_successfully');
                    vm.tableParams.reload();
                }

                function removeFail() {
                    logger.error('alert.daily_post_has_been_not_deleted_successfully');
                }
            }
        }
    }

})();
