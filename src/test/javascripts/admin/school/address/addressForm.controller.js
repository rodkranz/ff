(function () {
    'use strict';

    angular
        .module('app.school')
        .controller('SchoolAddressForm', SchoolAddressForm);

    SchoolAddressForm.$inject = [
        'logger', 'datacontext', '$state'
    ];

    /* @ngInject */
    function SchoolAddressForm(logger, datacontext, $state) {
        /* jshint validthis:true */
        var vm = this;

        // Variable
        vm.schoolAddress    = {};
        vm.errorsForm       = false;

        // Methods
        vm.doSave           = doSave;

        activate();
        ///////////

        /* Methods */
        function activate(){
            logger.debug('Admin > School > Address > Form > Loaded');
        }

        function doSave(school) {
            if (vm.formSchoolAddress.$invalid) {
                logger.error('please_fill_the_red_fields');
                vm.errorsForm = function () {
                    return vm.formSchoolAddress.$invalid
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
                logger.success('alert.address_has_been_saved_successfully');
                $state.go('grid.schoolList');
            }

            function createFail(result) {
                logger.error('alert.address_has_been_not_saved_successfully');
            }
        }
    }

})();
