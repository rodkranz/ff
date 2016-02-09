(function () {
    'use strict';

    angular
        .module('app.dailyPost')
        .controller('DailyPostForm', DailyPostForm);

    DailyPostForm.$inject = [
        'logger', 'dailyPost', 'datacontext', '$state'
    ];

    /* @ngInject */
    function DailyPostForm(logger, dailyPost, datacontext, $state) {
        /* jshint validthis:true */
        var vm = this;

        // Variable
        vm.dailyPost    = dailyPost || {};
        vm.errorsForm   = false;

        // Methods
        vm.doSave    = doSave;

        activate();
        ///////////

        /* Methods */
        function activate(){
            logger.debug('Admin > Daily Post List > Loaded');
        }

        function doSave(dailyPost) {
            if (vm.formDailyPost.$invalid) {
                logger.error('please_fill_the_red_fields');
                vm.errorsForm = function () {
                    return vm.formDailyPost.$invalid
                };
            }
            else
            {
                datacontext
                    .dailyPost
                    .saveDailyPost(dailyPost)
                    .then(createSuccess, createFail);
            }

            function createSuccess(result) {
                logger.success('alert.daily_post_has_been_saved_successfully');
                $state.go('grid.dailyPostList');
            }

            function createFail(result) {
                logger.error('alert.daily_post_has_been_not_saved_successfully');
            }
        }
    }

})();
