(function () {
    'use strict';

    angular
        .module('app.core')
        .factory('common', common);

    common.$inject = [
        '$q', '$rootScope', '$state', '$timeout', 'logger', 'config',
        '$window', '$filter', '$localStorage'
    ];

    /* @ngInject */
    function common($q, $rootScope, $state, $timeout, logger, config,
                    $window, $filter, $localStorage) {
        /* jshint validthis:true */

        var $win  = $($window);
        var $html = $('html');
        var $body = $('body');

        var services = {
                $broadcast          : $broadcast,
                $q                  : $q,
                $timeout            : $timeout,
                logger              : logger,
                config              : config,
                $win                : $($window),
                $window             : $window,
                $state              : $state,
                isMobile            : isMobile,
                isTouch             : isTouch,
                configLayout        : configLayout,
                translate           : translate,
                go                  : go,
                back                : back,
                $on                 : $rootScope.$on,
                formatNumber        : formatNumber,
                checkObjectEmpty    : checkObjectEmpty
            };

        return services;

        activate();
        ///////////////

        function activate() {
            $rootScope.$broadcast('loading', true);
        }

        function isMobile() {
            return $win.width() < config.deviceSizes.tablet;
        }

        function isTouch() {
            return $html.hasClass('touch');
        }

        function translate(text, attr) {
            return $filter('translate')(text, attr);
        }

        // Layout configure
        function configLayout(title, bool) {
            if (typeof(bool) !== 'undefined') {
                var layout = $localStorage.layout || {};
                layout[title] = bool;
                //$localStorage.layout[title] = bool;
                $localStorage.layout = layout;
            }
            else {
                return $localStorage.layout && $localStorage.layout[title];
            }
        }

        // Actions
        function $broadcast() {
            return $rootScope.$broadcast.apply($rootScope, arguments);
        }

        function go(url, params) {
            $state.go(url, params || {});
        }

        function back(n) {
            $window.history.back(n || -1);
        }

        function formatNumber(number, save, decimal) {
            try {
                if (!!save) {
                    return parseFloat(number);
                } else {
                    return $filter('number')(number, decimal || config.moneyPrecision);
                }
            } catch (e) {
                console.error(n, 'is not a number', e);
                return n;
            }
        }

        // End Layout methods
        function checkObjectEmpty(promiseCheck, callback) {
            var defer = $q.defer();

            promiseCheck
                .then(function (data) {
                    if (!!data && typeof data === 'object') {
                        return defer.resolve(data);
                    } else {
                        return defer.reject(data);
                    }
                })
                .catch(promiseCheck.reject);

            return defer.promise;
        }
    }
})();
