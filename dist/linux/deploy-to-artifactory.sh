#!/bin/bash
ArtifactId=workload-agent
Version=2.0-SNAPSHOT
GroupId=com.intel.isecl
MavenRepoPath=/root/.m2/repository/
cp out/${ArtifactId}-v2.*bin out/${ArtifactId}-${Version}.bin
mvn  install:install-file -Durl=file://${MavenRepoPath} -Dfile=out/${ArtifactId}-${Version}.bin -Dtype=bin -DartifactId=${ArtifactId} -DgroupId=${GroupId} -Dversion=${Version} -DpomFile=dist/linux/pom.xml -Dpackaging=bin
