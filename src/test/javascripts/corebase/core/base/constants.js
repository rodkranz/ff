(function () {
    'use strict';

    angular.module('app.core')
        .constant('toastr', toastr)
        .constant('Holder', Holder)
        .constant('swal',   swal);

})();
