(function () {
    'use strict';

    angular.module('app.core', [
        /*
         * Angular Modules
         */
        'ui.bootstrap', 'ngAnimate', 'ngStorage', 'ngCookies',
        'ngSanitize', 'ngStorage',

        /*
         * Ours Reusable cross app core repository
         */
        'blocks.logger', 'blocks.exception', 'blocks.interceptor',
        'blocks.router', 'blocks.auth',

        /*
         * Party Modules
         */
        'angularMoment', 'ui.router', 'pascalprecht.translate',
        'restangular', 'ui.bootstrap', 'ngTable'
    ]);

})();
