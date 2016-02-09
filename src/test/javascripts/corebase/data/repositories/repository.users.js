(function () {
    'use strict';

    angular
        .module('app.data')
        .factory('repository.user', user);

    user.$inject = ['$auth', '$rest'];

    /*@ngInject*/
    function user($auth, $rest) {
        /* jshint validthis:true */

        var repoName = "users";
        var services = {
            login: login,
            current: current
        };

        return services;

        //Login User
        function login(credentials) {
            return $rest
                .rest
                .all(repoName)
                .all('sign_in')
                .post(credentials)
                .then(loggin_success)
                .catch(loggin_failed);
        }

        //Renew User
        function current() {
            return $rest
                .restAuth
                .get('current');
        }


        // User Success
        function loggin_success(response) {
            $auth.addUser(response);
            $rest.getToken(true);
            return true;
        }

        // User Failed
        function loggin_failed(e) {
            return false;
        }
    }

})();
