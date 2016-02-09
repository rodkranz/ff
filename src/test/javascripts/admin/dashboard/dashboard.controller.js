(function () {
    'use strict';

    angular
        .module('app.dashboard')
        .controller('Dashboard', Dashboard);

    Dashboard.$inject = [];

    /* @ngInject */
    function Dashboard() {
        /* jshint validthis:true */
        var vm = this;

        activate();
        /* Methods */
        function activate(){
        }
    }

})();
