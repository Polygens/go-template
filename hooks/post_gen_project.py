#!/bin/env python3

import subprocess

import os
import shutil
import yaml

MANIFEST = "manifest.yml"

project_name = '{{ cookiecutter.project_name }}'
create_github_repository = '{{ cookiecutter.create_github_repository }}'

def delete_resources_for_disabled_features():
    with open(MANIFEST) as manifest_file:
        manifest = yaml.safe_load(manifest_file)
        for feature in manifest['features']:
            if not feature['param'] in feature['allowed']:
                print("removing resources for disabled feature {}...".format(feature['name']))
                for resource in feature['resources']:
                    delete_resource(resource)
    print("cleanup complete, removing manifest...")
    delete_resource(MANIFEST)


def delete_resource(resource):
    if os.path.isfile(resource):
        print("removing file: {}".format(resource))
        os.remove(resource)
    elif os.path.isdir(resource):
        print("removing directory: {}".format(resource))
        shutil.rmtree(resource)

def setup_go():
    subprocess.call(['go', 'get', '-u', 'golang.org/x/tools/cmd/goimports'])
    subprocess.call(['goimports', '-w', './'])

def setup_git():
    subprocess.call(['git', 'init'])
    subprocess.call(['git', 'add', '*'])
    subprocess.call(['git', 'reset', '--', 'github-template.json'])
    subprocess.call(['git', 'commit', '-m', 'initial commit'])
    subprocess.call(['git', 'tag', 'v0.1.0', '-m', 'first release'])

def setup_github():
    user_email = ""
    try:
        user_email = subprocess.check_output(['git', 'config', 'user.email']).strip()
    except subprocess.CalledProcessError:
        print("Failed to get user email")
        os._exit(os.EX_NOTFOUND) 

    subprocess.call(['curl', '-X', 'POST', '-u', user_email, '-H', 'Content-Type: application/json', '-d', '@./github-template.json', 'https://api.github.com/user/repos'])

    subprocess.call(['rm', './github-template.json'])
    subprocess.call(['git', 'remote', 'add', 'origin', 'https://github.com/Polygens/' + project_name + '.git'])
    subprocess.call(['git', 'push', '-u', 'origin', 'master'])
    subprocess.call(['git', 'push', 'origin', '--tags'])

if __name__ == "__main__":
    delete_resources_for_disabled_features()
    setup_go()
    setup_git()
    if create_github_repository == "yes":
        setup_github()
