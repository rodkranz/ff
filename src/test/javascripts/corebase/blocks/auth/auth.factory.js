(function () {
    'use strict';

    angular
        .module('blocks.auth')
        .factory('$auth', auth);

    auth.$inject = ['$window', '$rootScope'];

    /*@ngInject*/
    function auth($window, $rootScope) {
        /* jshint validthis:true */

        var service = {
            logOut: logOut,
            isAuthenticated: isAuthenticated,
            getToken: getToken,
            addUser: addUser,
            getUser: getUser,
            updateStatus: updateStatus,
            hasRole: hasRole
        };

        return service;

        function logOut() {
            if (!isAuthenticated()) {
                return false;
            }
            $window.localStorage.removeItem('currentUser');
            updateStatus();

            return true;
        }

        function isAuthenticated() {
            return !!$window.localStorage.currentUser;
        }

        function getToken() {
            if (!isAuthenticated()) {
                return '';
            }

            var user = angular.fromJson($window.localStorage.currentUser);
            return user.authentication_token;
        }

        function addUser(user) {
            if (!user) {
                return false;
            }
            $window.localStorage.currentUser = angular.toJson(user);
            updateStatus();
            return true;
        }

        function getUser(field) {
            if (!isAuthenticated()) {
                return false;
            }

            try {
                var user = angular.fromJson($window.localStorage.currentUser);
                if (!!field && !!user && user.hasOwnProperty(field)) {
                    return user[field];
                }
                else {
                    return user;
                }
            } catch (e) {
                console.error(e);
                return false;
            }
        }

        function hasRole(role) {
            return getUser('role') == role;
        }

        function updateStatus() {
            $rootScope.isLogged = isAuthenticated() || false;
            $rootScope.user     = getUser()         || {};
            $rootScope.$broadcast('auth:updated');
        }
    }
})();
