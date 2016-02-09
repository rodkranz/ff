(function () {
    'use strict';

    angular
        .module('app.school')
        .controller('SchoolGalleryForm', SchoolGalleryForm);

    SchoolGalleryForm.$inject = [
        'logger', 'datacontext', '$state'
    ];

    /* @ngInject */
    function SchoolGalleryForm(logger, datacontext, $state) {
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
            logger.debug('Admin > School > Gallery > Form > Loaded');
        }

        function doSave(school) {
            if (vm.formSchoolGallery.$invalid) {
                logger.error('please_fill_the_red_fields');
                vm.errorsForm = function () {
                    return vm.formSchoolGallery.$invalid
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
                logger.success('alert.gallery_has_been_saved_successfully');
                $state.go('grid.schoolList');
            }

            function createFail(result) {
                logger.error('alert.gallery_has_been_not_saved_successfully');
            }
        }
    }

})();
