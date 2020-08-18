import accessConversionTable

def CreateUser(clientGitlab, email, password, username, name):
    user = clientGitlab.client.users.create({'email': email,
                        'password': password,
                        'username': username,
                        'name': name})
    #ERROR CHECKING TODO
    return user, True


def getGitlabAccess(user):
    job=user.role
    access_level = accessConversionTable.getGitlabConversionTable(job)
    return access_level