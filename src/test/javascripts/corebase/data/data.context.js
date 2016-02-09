(function () {
    'use strict';

    angular
        .module('app.data')
        .factory('datacontext', datacontext);

    datacontext.$inject = ['$injector', '$q', 'exception', 'config'];

    /*@ngInject*/
    function datacontext($injector, $q, exception, config) {
        /* jshint validthis:true */

        var isPrimed = false;
        var primePromise;
        var repoNames = config.repositories;

        var service = {
            ready: ready
        };

        activate();
        return service;

        function activate() {
            defineLazyLoadedRepos();
        }

        function defineLazyLoadedRepos() {
            repoNames.forEach(function (name) {
                Object.defineProperty(service, name, {
                    get: function () {
                        var repo = getRepo(name);
                        return repo;
                    }
                });
            })
        }

        function getRepo(repoName) {
            var fullRepoName = 'repository.' + repoName;
            var factory = $injector.get(fullRepoName);
            return factory;
        }

        function prime() {
            if (primePromise) {
                return primePromise;
            }

            primePromise = $q.when(true).then(success);
            return primePromise;

            function success() {
                isPrimed = true;
            }
        }

        function ready(nextPromises) {
            var readyPromise = primePromise || prime();

            return readyPromise
                .then(function () {
                    return $q.all(nextPromises);
                })
                .catch(exception.catcher("'ready' function failed"));
        }

    }

})();
