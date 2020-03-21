# Pekka
Deploy and manage multiple wordpress sites with traefik and docker with Let's encrypt support. 

Pekka is a simple wrapper over `docker-compose`. Pekka generates and uses normal docker-compose files.  

# Prerequisites
docker and docker-compose should be installed and configured such that the user running `pekka` should have access and permissions to run `docker` and `docker-compose` commands

# Usage
1. Initialize pekka with the ``init`` command. 
  
   This creates the required files in the current directory.

    ```
    $ pekka init                     
    Enter traefik dashboard URL: <traefik dashboard URL>
    Enter let's encrypt email: <acme email>
    Creating network "pekkatraefik_webgateway" with driver "bridge"
    Creating pekkatraefik_proxy_1 ... 
    Creating pekkatraefik_proxy_1 ... done
    ```

    `<traefik dashboard URL>` is the URL where the traefik server's dashboard will be exposed.

    <acme email> is used for lets encrypt configurations.

2. Create a wordpress deployment using the ``create`` command

    ```
    $ pekka create toys
    Enter domain name: example.com
    Add entry for www.example.com? y
    Creating network "toys_default" with the default driver
    Pulling wordpress (nithinbose/wordpress:latest)...
    latest: Pulling from nithinbose/wordpress
    aa18ad1a0d33: Pull complete
    29d5f85af454: Pull complete
    eca642e7826b: Pull complete
    3638d91a9039: Pull complete
    3646a95ab677: Pull complete
    628b8373e193: Pull complete
    c24a2b2280ed: Pull complete
    f968b84cbbbc: Pull complete
    60fafe14064c: Pull complete
    bac57a95ddf1: Pull complete
    056ffd8ba0fc: Pull complete
    b595ac5a4e55: Pull complete
    5b72115923ec: Pull complete
    81b6cd799f34: Pull complete
    83faafba8a33: Pull complete
    577a4001244f: Pull complete
    69765c2499ed: Pull complete
    0044a72ca220: Pull complete
    5481d2b46462: Pull complete
    fcab5f51b65c: Pull complete
    0de0045cbc4b: Pull complete
    Digest: sha256:0f00bc21638db44478039e70e56ba40a0835b034a05300a4dcbfce2f86e26495
    Status: Downloaded newer image for nithinbose/wordpress:latest
    Creating toys_mysql_1 ... 
    Creating toys_mysql_1 ... done
    Creating toys_wordpress_1 ... 
    Creating toys_wordpress_1 ... done
    ```

    Pekka uses a custom wordpress docker image image ``nithinbose/wordpress``hosted on docker hub.

3. Remove a wordpress deployment using the ``remove`` command

    ```
    $ pekka remove toys
    Stopping toys_wordpress_1 ... done
    Stopping toys_mariadb_1 ... done
    Deployment stopped
    Going to remove toys_wordpress_1, toys_mariadb_1
    Removing toys_wordpress_1 ... done
    Removing toys_mariadb_1 ... done
    Removing files...
    Deployment removed
    ```

4. Update a wordpress deployment using the ``update`` command

    ```
    $ ./pekka update test

    Pulling updates...
    Pulling mariadb (mariadb:latest)...
    latest: Pulling from library/mariadb
    aa18ad1a0d33: Already exists
    fdb8d83dece3: Already exists
    75b6ce7b50d3: Already exists
    ed1d0a3a64e4: Already exists
    b153f26fb6de: Already exists
    7df63694312a: Already exists
    8efb9b48cc94: Already exists
    8c9580569876: Pull complete
    e3ab27d2678e: Pull complete
    7ba7e552e2cd: Pull complete
    3ececfe301de: Pull complete
    9190ef68b220: Pull complete
    Digest: sha256:97c90d8672995c9471d374ff1ef43eba14197eda9fcafaeb2d007be5838fccbd
    Status: Downloaded newer image for mariadb:latest
    Pulling wordpress (nithinbose/wordpress:latest)...
    latest: Pulling from nithinbose/wordpress
    aa18ad1a0d33: Already exists
    29d5f85af454: Already exists
    eca642e7826b: Already exists
    3638d91a9039: Already exists
    3646a95ab677: Already exists
    628b8373e193: Already exists
    c24a2b2280ed: Already exists
    f968b84cbbbc: Already exists
    60fafe14064c: Already exists
    bac57a95ddf1: Already exists
    056ffd8ba0fc: Already exists
    3c7a6d81f935: Pull complete
    1538d9314280: Pull complete
    6226f413dc4f: Pull complete
    c85b972ccc30: Pull complete
    c055dd59d170: Pull complete
    53a5c17cf43e: Pull complete
    f1ff0a4875c3: Pull complete
    d773f2dd14c8: Pull complete
    303b2c36f914: Pull complete
    a02bfbd68f98: Pull complete
    Digest: sha256:31150ce4e433b57a07fbd256e79853d08851d830aee6ced7cc40cd858e69e961
    Status: Downloaded newer image for nithinbose/wordpress:latest
    Updates pulled

    Stopping current deployment...
    Stopping test_wordpress_1 ... done
    Stopping test_mariadb_1 ... done
    Going to remove test_wordpress_1, test_mariadb_1
    Removing test_wordpress_1 ... done
    Removing test_mariadb_1 ... done

    Restarting deployment with updates...
    Creating test_mariadb_1 ...
    Creating test_mariadb_1 ... done
    Creating test_wordpress_1 ...
    Creating test_wordpress_1 ... done

    Deployment updated
    ```