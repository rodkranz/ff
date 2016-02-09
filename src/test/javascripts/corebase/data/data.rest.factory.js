(function () {
    'use strict';

    angular.module('app.data')
        .factory('$rest', RestFul);

    RestFul.$inject = ['Restangular', '$rootScope', '$auth'];

    /* @ngInject */
    function RestFul(Restangular, $rootScope, $auth) {
        /* jshint validthis:true */
        var self = this;

        var services = {
            getToken:    getToken,
            rest:        Restangular,
            restAuth:    Restangular.withConfig(RestangularAuth),
            restFul:     Restangular.withConfig(RestangularFull),
            restFulAuth: Restangular.withConfig(RestangularFullAuth)
        };

        return services;
        ////////////////

        function getToken(force) {
            if (force === true) {
                self.myToken = {'auth_token': $auth.getToken()};
            }
            return self.myToken;
        }

        function RestangularAuth(RestangularConfigurer) {
            $rootScope.$on('auth:updated', function () {
                RestangularConfigurer.setDefaultRequestParams(getToken(true));
            });
            RestangularConfigurer.setDefaultRequestParams(getToken(true));
            setDefaultRequestParams(RestangularConfigurer);
        }

        function RestangularFull(RestangularConfigurer) {
            RestangularConfigurer.setFullResponse(true);
        }

        function RestangularFullAuth(RestangularConfigurer) {
            // Fix problem when changed the user token
            $rootScope.$on('auth:updated', function () {
                RestangularConfigurer.setDefaultRequestParams(getToken(true));
            });

            RestangularConfigurer.setDefaultRequestParams(getToken(true));
            RestangularConfigurer.setFullResponse(true);
        }

        function setDefaultRequestParams(RestangularConfigurer){
            RestangularConfigurer.setDefaultRequestParams(getToken(true));
        }
    }
})();
