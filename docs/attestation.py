from enum import Enum
from datetime import datetime
    
class Action_Type(Enum):
  consent_revoked = "consent revoked"
  consent_accepted = "consent accepted"
  consent_denied = "consent denied"
  authorization_granted = "authorization granted"
  data_subject_request_make_request = "data subject request: initiate"
  data_use = "data use"
  

class Action:
  def __init__(self, type: Action_Type, information: JSON) -> None:
    self.type = type
    self.information = information
    
class Party:
  def __init__(self, name: str, data_controller: Set({"consumer", "data_provider", "data_recipient"})) -> None:
    self.name = name
    self.data_controller = data_controller
  
class Attestation:
  def __init__(self, party: Party, action: Action, timestamp: datetime) -> None:
    self._id = _id
    self.party = party
    self.action = action
    self.timestamp = timestamp
      

class AttestationRecords:
  def __init__(self, party: Party) -> None:
    self.party = party
    self.attestations = [] # list of instances of the Attestation class
    
  def attest(self, action: Action) -> Attestation:
    new_attestation = Attestation(self.party, action, datetime.now())
    self.attestations.append(new_attestation)
    return new_attestation
  
  def all(self) -> List[Attestation]:
    return self.attestations
  
  # generate an object saying that up_to_date_attestations includes all actions between timestamp1 and timestamp2
  def up_to_date(self, timestamp1: datetime, timestamp2: datetime) -> List[Attestation]:
    up_to_date_attestations = [attestation for attestation in self.attestations if d1 <= attestation.timestamp <= d2]
		return up_to_date_attestations
