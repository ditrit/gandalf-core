from gitlab import Gitlab
import pygitlab.template

def gitlabRecognize(u):
    return True
    
def getGitlabAccess(u):
    return 40

def CreateProject(clientGitlab, name, team, template_name = None):
    if template_name!= None :
        project = clientGitlab.client.projects.create({'name': name, 'template_name': template_name})
    else: 
        project = clientGitlab.client.projects.create({'name': name})
    #TODO : création aussi du "team account"??
    members = []
    for u in team: #on définit ici la team comme une liste des emails des membres
        #TODO fonction gitlabRecognize et définir attribut de user
        if gitlabRecognize(u):
            #requette SQL si user est dans une table de donnée??
            userGitlab=clientGitlab.client.users.list(email = u)[0]
            members.append( project.members.create({'user_id': userGitlab.id, 'access_level': getGitlabAccess(u) })) 
        else : 
            new_user = user.CreateUser(clientGitlab, u.email, "0000", u.username, u.name) #mot de passe initialisé à 0000
            member = project.members.create({'user_id': new_user.id, 'access_level':
                                 getGitlabAccess(u)})
            members.append( member ) 
            #TODO
            print("The user", u, "is not known by Gitlab. The user", member, " has been created instead"  )
    
    #TODO ERROR CHECKING 
    return [project.id , True]


def AddMember(clientGitlab, project_id, user_email):
    project = clientGitlab.client.projects.get(project_id)
     #TODO fonction gitlabRecognize et définir attribut de user
    listUsers=clientGitlab.client.users.list(email = user_email)
    if len(listUser)!= 0 and listUser[0].email==user_email :
        userGitlab=clientGitlab.client.users.list(email = user_email)[0]
        member = project.members.create({'user_id': userGitlab.id, 'access_level': getGitlabAccess(u) })
    else : 
        new_user = user.CreateUser(clientGitlab, user.email, "0000", user.username, user.name) #mot de passe initialisé à 0000
        member = project.members.create({'user_id': new_user.id, 'access_level':
                                 getGitlabAccess(user_email)}) 
        #TODO
        print("The user", user_email, "is not known by Gitlab. The user", new_user, " has been created instead"  )
    
    #TODO ERROR CHECKING 
    return True

def RemoveMember(clientGitlab, project_id, user_email):
    project = clientGitlab.client.projects.get(project_id)
    userGitlab=clientGitlab.client.users.list(email=  user_email)[0]
    #TODO error if user is not part of the project
    project.members.delete(userGitlab.id)
    #TODO ERROR CHECKING 
    return True

