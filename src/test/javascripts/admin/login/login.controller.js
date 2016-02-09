(function () {
    'use strict';

    angular
        .module('app.login')
        .controller('Login', Login);

    Login.$inject = [
        'datacontext', '$state', '$location', '$rootScope'
    ];

    /* @ngInject */
    function Login(datacontext, $state, $location, $rootScope)
    {
        /* jshint validthis:true */
        var vm = this;

        vm.credentials = {};

        vm.submitted  = false;
        vm.doLogin    = doLogin;
        vm.interacted = interacted;

        activate();
        ///////////////////////////

        function activate()
        {
            $('[name=uEmail]').focus();
        }

        function doLogin(credentials)
        {
            vm.submitted = true;
            datacontext
                .user
                .login({user: credentials})
                .then(completed);

            function completed(userLogged)
            {
                if (userLogged)
                {
                    $rootScope.isLogged = true;
                    $state.go('grid.dashboard');
                }
                else
                {
                    vm.credentials.password = '';
                    $('[name=uPassword]').focus();
                }
            }
        }

        function interacted(field)
        {
            return vm.submitted || field.$dirty;
        }
    }

})();


