(function () {
    'use strict';

    angular
        .module('app.school')
        .controller('SchoolCertificateForm', SchoolCertificateForm);

    SchoolCertificateForm.$inject = [
        'logger', 'datacontext', '$state'
    ];

    /* @ngInject */
    function SchoolCertificateForm(logger, datacontext, $state) {
        /* jshint validthis:true */
        var vm = this;

        // Variable
        vm.errorsForm   = false;

        // Methods
        vm.doSave    = doSave;

        activate();
        ///////////

        /* Methods */
        function activate(){
            logger.debug('Admin > School > Certificate > Form > Loaded');
        }

        function doSave(school) {
            if (vm.formSchoolCertificate.$invalid) {
                logger.error('please_fill_the_red_fields');
                vm.errorsForm = function () {
                    return vm.formSchoolCertificate.$invalid
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
                logger.success('alert.certificate_has_been_saved_successfully');
                $state.go('grid.schoolList');
            }

            function createFail(result) {
                logger.error('alert.certificate_has_been_not_saved_successfully');
            }
        }
    }

})();
