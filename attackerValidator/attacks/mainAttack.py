from .ssti import NumberSSTI

# Wrapper over all attacks
# Please ensure all of your attacks are called in this function
# Should return true if no attacks were successful, and false if any one attack succeded.
# Please return a string with error message if an error is caught - it will not be show to the participants
def AttackWrapper(VulnMachineLocation):
    # Number based SSTI attack (aka. indicator)
    if NumberSSTI(VulnMachineLocation):
        return False

    # TODO: Add more tests
    
    # Otherwise...
    return True