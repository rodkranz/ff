(function () {
    'use strict';

    angular
        .module('app.school')
        .controller('SchoolForm', SchoolForm);

    SchoolForm.$inject = [
        'logger', 'school', 'datacontext', '$state'
    ];

    /* @ngInject */
    function SchoolForm(logger, school, datacontext, $state) {
        /* jshint validthis:true */
        var vm = this;

        // Variable
        vm.school    = school || {};
        vm.errorsForm   = false;

        // Methods
        vm.doSave    = doSave;

        activate();
        ///////////

        /* Methods */
        function activate(){
            logger.debug('Admin > School > List > Loaded');
        }

        function doSave(school) {
            if (vm.formSchool.$invalid) {
                logger.error('please_fill_the_red_fields');
                vm.errorsForm = function () {
                    return vm.formSchool.$invalid
                };
            }
            else
            {
                datacontext
                    .school
                    .saveSchool(school)
                    .then(createSuccess, createFail);
            }

            function createSuccess(result) {
                logger.success('alert.School_has_been_saved_successfully');
                $state.go('grid.schoolList');
            }

            function createFail(result) {
                logger.error('alert.School_has_been_not_saved_successfully');
            }
        }
    }

})();
