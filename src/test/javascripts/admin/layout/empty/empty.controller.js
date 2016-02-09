(function () {
    'use strict';

    angular
        .module('app.layout')
        .controller('Empty', Empty);

    Empty.$inject = ['$rootScope'];

    /*@ngInject*/
    function Empty($rootScope) {
        /* jshint validthis:true */
        var vm = this;
    }

})();
