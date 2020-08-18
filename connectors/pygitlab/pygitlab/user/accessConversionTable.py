#Conversion table for Gitlab access 
#attention à différencier des acces au fct des workers, ou pas?

gitlabConversionTable=dict()
access_levels={}
access_levels["No access"] = 0
access_levels["Guest"] = 10
access_levels["Reporter"] = 20
access_levels["Developer"] = 30
access_levels["Maintainer"] = 40
access_levels["Owner"] = 50
#ajouter les sendCommand??

job=["Scrum Master", "Development", "Product owner", "Client", "Management", "User"]

#Type DevOps
gitlabConversionTable["Team account"] = "Owner"
gitlabConversionTable["Scrum Master"] = "Maintainer"
gitlabConversionTable["Development"] = "Developer"
gitlabConversionTable["Product owner"] = "Owner"
gitlabConversionTable["Client"] = "Reporter"
gitlabConversionTable["Management"] = "Maintainer"
gitlabConversionTable["User"] = "Guest"

#type classique
gitlabConversionTable["Administrateur"] = "Maintainer"
gitlabConversionTable["Opérateur"] = "Developer"


def getGitlabConversionTable(job):
    return access_levels[gitlabConversionTable[job]]
