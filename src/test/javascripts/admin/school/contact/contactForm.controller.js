(function () {
    'use strict';

    angular
        .module('app.school')
        .controller('SchoolContactForm', SchoolContactForm);

    SchoolContactForm.$inject = [
        'logger', 'datacontext', '$state'
    ];

    /* @ngInject */
    function SchoolContactForm(logger, datacontext, $state) {
        /* jshint validthis:true */
        var vm = this;

        // Variable
        vm.schoolContact = {};
        vm.errorsForm    = false;

        // Methods
        vm.doSave        = doSave;

        activate();
        ///////////

        /* Methods */
        function activate(){
            logger.debug('Admin > School > Contact > Form > Loaded');
        }

        function doSave(school) {
            if (vm.formSchoolContact.$invalid) {
                logger.error('please_fill_the_red_fields');
                vm.errorsForm = function () {
                    return vm.formSchoolContact.$invalid
                };
            }
            else
            {
                //datacontext
                //    .school
                //    .saveSchool(school)
                //    .then(createSuccess, createFail);
                createSuccess();
            }

            function createSuccess(result) {
                logger.success('alert.contact_has_been_saved_successfully');
                $state.go('grid.schoolList');
            }

            function createFail(result) {
                logger.error('alert.contact_has_been_not_saved_successfully');
            }
        }
    }

})();
