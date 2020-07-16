def CreateIssue(githubClient, repositoryName, title, body):

    repository = githubClient.client.get_repo(repositoryName)

    repository.create_issue(title=title, body = body)
    #ERROR CHECKING TODO

    return True
