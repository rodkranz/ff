(function () {
    'use strict';

    angular
        .module('app.layout')
        .controller('Navbar', Navbar);

    Navbar.$inject = ['$rootScope', '$auth', '$state'];

    /*@ngInject*/
    function Navbar($rootScope, $auth, $state) {
        /* jshint validthis:true */
        var vm = this;

        vm.signOut = signOut;
        activate();

        function activate()
        {
            updateLogin();
            $rootScope.$on('auth:updated', updateLogin);
        }

        function updateLogin()
        {
            $rootScope.isLogged = $auth.isAuthenticated();
        }

        function signOut()
        {
            $auth.logOut();
            $state.go('empty.login');
        }
    }
})();
