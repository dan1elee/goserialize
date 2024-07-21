from enum import Enum, auto
class WrongFormErrorType(Enum):
    LengthTooShort = 0
    LengthNotEqual = auto()
    TypeNotSupport = auto()

class WrongFormException(Exception):
    def __init__(self,  errorType: WrongFormErrorType):
        self.errorType = errorType
        super().__init__("unserialize from wrong form")
    
    def __str__(self):
        return f"WRONG FORM EXCEPTION (type: {self.errorType.name})"
    
