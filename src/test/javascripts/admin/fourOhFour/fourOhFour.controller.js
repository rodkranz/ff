(function () {
    'use strict';

    angular
        .module('app.fourOhFour')
        .controller('FourOhFour', FourOhFour);

    FourOhFour.$inject = [];

    /* @ngInject */
    function FourOhFour() {
        /* jshint validthis:true */
        var vm = this;

        activate();
        /* Methods */
        function activate(){
        }
    }

})();
