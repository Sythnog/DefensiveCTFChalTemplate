import requests

# Number based
def NumberSSTI(Location):
    exploit = "{{ 7*7 }}"
    req = requests.get(Location + exploit, timeout=5)
    # TODO: Make the check harder to bypass without actually solving the challenge
    if req.text == "<h1>49 was not found :(</h1>":
        return True
    else:
        return False

# For debug purposes, run if called directly
if __name__ == "__main__":
    print(NumberSSTI("http://172.17.0.3/"))
