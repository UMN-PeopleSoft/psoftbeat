# psoftbeat

## Summary
Psoftbeat is a beat based on metricbeat which was generated with metricbeat/metricset generator.
Psoftbeat captures metrics for PeopleSoft targets like Weblogic, App Server (Tuxedo) and Scheduler(Tuxedo).

The psoftbeat monitor uses the [psoftjmx](https://github.com/UMN-PeopleSoft/psoftjmx) library that provides a fast concurrent loader of JMX metrics from all PeopleSoft target types.

See releases for current packaged deployment ready to be used.

Steps for running the beat are on the wiki page https://github.com/UMN-PeopleSoft/psoftbeat/wiki/Implementing-the-psoftbeat-on-the-ElasticSearch-and-Kibana



## Development 

## Getting started

To get started run the following command. This command should only be run once.

```
make setup
```

It will ask you for the module and metricset name. Insert the name accordingly.

To compile your beat run `make`. Then you can run the following command to see the first output:

```
psoftbeat -e -d "*"
```

In case further modules are metricsets should be added, run:

```
make create-metricset
```

After updates to the fields or config files, always run

```
make collect
```

This updates all fields and docs with the most recent changes.

## Use vendoring

We recommend to use vendoring for your beat. This means the dependencies are put into your beat folder. The beats team currently uses [govendor](https://github.com/kardianos/govendor) for vendoring.

```
govendor init
govendor update +e
```

This will create a directory `vendor` inside your repository. To make sure all dependencies for the Makefile commands are loaded from the vendor directory, find the following line in your Makefile:

```
ES_BEATS=${GOPATH}/src/github.com/elastic/beats
```

Replace it with:
```
ES_BEATS=./vendor/github.com/elastic/beats
```


## Versioning

We recommend to version your repository with git and make it available on Github so others can also use your project. The initialise the git repository and add the first commits, you can use the following commands:

```
git init
git add README.md CONTRIBUTING.md
git commit -m "Initial commit"
git add LICENSE
git commit -m "Add the LICENSE"
git add .gitignore
git commit -m "Add git settings"
git add .
git reset -- .travis.yml
git commit -m "Add psoftbeat"
```

## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.
