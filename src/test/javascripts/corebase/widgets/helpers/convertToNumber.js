(function () {
    'use strict';

    angular
        .module('app.widgets')
        .directive('convertToNumber', convertToNumber);

    convertToNumber.$inject = [];

    /* @ngInject */
    function convertToNumber() {
        return {
            require: 'ngModel',
            link: function(scope, el, attr, ctrl) {
                ctrl.$parsers.push(function(value) {
                    return parseInt(value, 10);
                });

                ctrl.$formatters.push(function(value) {
                    return value.toString();
                });
            }
        }
    }
})();