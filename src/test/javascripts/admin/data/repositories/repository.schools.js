(function () {
    'use strict';

    angular
        .module('app.data')
        .factory('repository.school', school);

    school.$inject = ['$rest'];

    /*@ngInject*/
    function school($rest) {
        /* jshint validthis:true */

        var resourceName         = "schools";
        var resourceServices     = {
            findAllSchools:         findAllSchools,
            findSchoolById:         findSchoolById,
            saveSchool:             saveSchool,
            removeSchool:           removeSchool,
            updateSchoolAddress:    updateSchoolAddress,
            updateSchoolContact:    updateSchoolContact
        };

        return resourceServices;

        function findAllSchools(params) {
            return $rest
                .restFulAuth
                .all(resourceName)
                .getList(params);
        }

        function findSchoolById(id) {
            return $rest
                .restAuth
                .one(resourceName, id)
                .get();
        }

        function saveSchool(school) {
            if ( school.restangularized ) {
                return school.save();
            } else {
                return $rest
                    .restFulAuth
                    .all(resourceName)
                    .post(school);
            }
        }

        function removeSchool(id) {
            return $rest
                .restFulAuth
                .one(resourceName, id)
                .remove();
        }

        function updateSchoolAddress(address) {
        }

        function updateSchoolContact(contact) {
        }
      
    }

})();