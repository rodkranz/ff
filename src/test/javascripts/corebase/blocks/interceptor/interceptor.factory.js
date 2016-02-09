(function () {
    'use strict';

    angular
        .module('blocks.interceptor')
        .run(Interceptor);

    Interceptor.$inject = [
        'logger', 'Restangular', '$location', 'config'
    ];

    /*@ngInject*/
    function Interceptor(logger, Restangular, $location, config) {

        Restangular.setErrorInterceptor(errorInterceptor);

        // Interceptor if there is any error.
        function errorInterceptor(response, deferred, responseHandler) {
            var message;
            //logger.debug( response.data, null, response.status );
            switch (response.status) {
                case 401:
                    message = response.data && response.data.error || response.statusText;
                    logger.error(message);
                    $location.path(config.pathDenied);
                    return false;
                    break;
                case 400:
                    message = response.data && response.data.message;
                    logger.warning('message.' + message);
                    break;
                case 404:
                    logger.error('message.page_not_found', undefined, '404');
                    $location.path(config.pathNotFound);
                    break;
                case 500:
                    var extra;
                    message = response.data && response.data.meta && response.data.meta.message || 'Erro internal';
                    logger.error(message, response.status, extra);
                    break;
            }
            return true;

        }
    }
})();