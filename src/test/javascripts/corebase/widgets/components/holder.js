(function () {
    'use strict';

    angular
        .module('app.widgets')
        .directive('holderFix', holderFix);

    holderFix.$inject = ['$timeout'];

    /* @ngInject */
    function holderFix($timeout) {
        return {
            link: function (scope, element, attrs) {
                $timeout(function () {
                    Holder.run({images: element[0], nocss: true});
                }, 100);
            }
        }
    }
})();