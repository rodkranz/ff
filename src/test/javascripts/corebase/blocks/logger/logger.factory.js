(function () {
    'use strict';

    angular
        .module('blocks.logger')
        .factory('logger', Logger);

    Logger.$inject = [
        '$log', 'toastr', '$filter'
    ];

    /* @ngInject */
    function Logger($log, toastr, $filter) {
        /* jshint validthis:true */
        return {
            error:   error,
            info:    info,
            success: success,
            warning: warning,
            debug:   debug
        };

        function translate(text) {
            return $filter('translate')(text);
        }

        function error(message, data, title) {
            toastr.error(translate(message), translate(title));
            $log.error('Error: ' + translate(message), translate(data));
        }

        function info(message, data, title) {
            toastr.info(translate(message), translate(title));
            $log.info('Info: ' + translate(message), translate(data));
        }

        function success(message, data, title) {
            toastr.success(translate(message), translate(title));
            $log.debug('Success: ' + translate(message), translate(data));
        }

        function warning(message, data, title) {
            toastr.warning(translate(message), translate(title));
            $log.warn('Warn: ' + translate(message), translate(data));
        }

        function debug(message, data, title) {
            if (typeof message == 'object') {
                var bMessage = JSON.stringify(message);
                $log.warn(message, translate(title));
                //toastr.info(bMessage, translate(title));
            } else {
                //toastr.info(translate(message), translate(title));
                $log.warn('Debug: ' + translate(message), translate(data));
            }
        }
    }
})();
