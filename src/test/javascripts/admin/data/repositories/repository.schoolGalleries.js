(function () {
    'use strict';

    angular
        .module('app.data')
        .factory('repository.schoolGalleries', schoolGalleries);

    schoolGalleries.$inject = ['$rest'];

    /*@ngInject*/
    function schoolGalleries($rest) {
        /* jshint validthis:true */

        var resourceName         = "schools";
        var subResourceName      = "gallery";
        var resourceServices     = {
            saveGalleryBySchool:             saveGalleryBySchool,
            removeGalleryBySchool:           removeGalleryBySchool,
        };

        return resourceServices;

        function saveGalleryBySchool(schoolId, gallery) {
            if ( gallery.restangularized ) {
                return gallery.save();
            } else {
                return $rest
                    .restFulAuth
                    .one(resourceName, schoolId)
                    .all(subResourceName)
                    .post(gallery);
            }
        }

        function removeGalleryBySchool(schoolId, galleryId) {
            return $rest
                .restFulAuth
                .one(resourceName, schoolId)
                .one(subResourceName, galleryId)
                .remove();
        }      
    }

})();