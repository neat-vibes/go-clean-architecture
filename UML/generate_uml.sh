#!/bin/bash

goplantuml -recursive -aggregate-private-members -show-aggregations -show-compositions -show-connection-labels -show-implementations /root/CleanInfrastructure/go-clean-architecture/ > diagram_file_name.puml