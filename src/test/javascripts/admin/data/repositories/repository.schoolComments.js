(function () {
    'use strict';

    angular
        .module('app.data')
        .factory('repository.schoolComments', schoolComments);

    schoolComments.$inject = ['$rest'];

    /*@ngInject*/
    function schoolComments($rest) {
        /* jshint validthis:true */

        var resourceName         = "schools";
        var subResourceName      = "comments";

        var resourceServices     = {
            findAllComments:   findAllComments,
            findCommentById:   findCommentById,
            saveComment:       saveComment,
            removeComment:     removeComment
        };

        return resourceServices;

        function findAllComments(schoolId, params) {
            return $rest
                .restFulAuth
                .one(resourceName, schoolId)
                .all(subResourceName)
                .getList(params);
        }

        function findCommentById(schoolId, id) {
            return $rest
                .restAuth
                .one(resourceName, schoolId)
                .one(resourceName, id)
                .get();
        }

        function saveComment(schoolId, comment) {
            if ( comment.restangularized ) {
                return comment.save();
            } else {
                return $rest
                    .restFulAuth
                    .one(resourceName, schoolId)
                    .all(subResourceName)
                    .post(comment);
            }
        }

        function removeComment(schoolId, commentId) {
            return $rest
                .restFulAuth
                .one(resourceName, schoolId)
                .one(subResourceNames, commentId)
                .remove();
        }
       
    }

})();