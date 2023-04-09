#!/usr/bin/env bash

PASSWORD="F@nfuryG#n3sis@fury"
GAS_PRICES="0.000025utfury"
CHAIN_ID="redshift-7"
NODE="https://rpc.furya.xyz:26657"

USERS=(FuryGuardian-1 FuryGuardian-2 genArg genBra genBro genBuf genInd genNY genLA genPSG genSF genStReserve)
MNEMONICS=(
	'dismiss surround benefit alarm crew hood spoil walnut photo promote special champion short uncover time primary core emotion hill lizard ocean room online together'
	'neglect lucky term swallow rotate mask runway voyage bridge female guilt round panic episode outdoor good below climb high method cushion donor laptop dragon'
	'noble door shed drift melt catalog mention blush kingdom sheriff churn congress cover danger use blame cloud corn put worth opinion move bulk oyster'
	'cotton mom coast curious degree green demand place margin express swarm strategy stomach suit fiscal luxury accuse industry horn tortoise ramp neglect noise infant'
	'valid spike survey stick traffic upper multiply two coyote desert situate twin track foam method inmate survey furnace ugly general engage exist correct seven'
	'release cushion enable cruise about dutch lazy pond desert sick curious run tribe autumn pulp jaguar hamster result where gravity rich trick foam blanket'
	'camp ability once yard survey inmate rescue chief clay legend paper echo sadness rebel kitchen accident nut close tell monitor barrel rent letter paddle'
	'dutch resource claim cliff choose rib path math vibrant drive eyebrow lunar travel kingdom flag crop height album crowd wife embody kitchen remove solve'
	'sting curve flight gauge rough egg frequent foot toss expire search horse dignity man sketch cabin flower extra dentist because room smile sort surround'
	'pill mention excite pact audit garment obey claw ice play impulse spirit anger acquire elbow oven office enter hockey often smart save intact gap'
	'remind latin happy coffee pizza undo apology team crane never glow daughter color lift disease energy abandon wrist payment seed box ankle average hurdle'
	'spare slow hat sign torch hill scheme carpet tuition swap auction ride dry smooth will cook mandate ten unveil insect december manual wave twelve'
	)

wait_chain_start() {
  RET=$(xfury status 2>&1)
  if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
    while :; do
      RET=$(xfury status 2>&1)
      if [[ ($RET == Error*) || ($RET == *'"latest_block_height":"0"'*) ]]; then
        sleep 1
      else
        echo "A few more seconds..."
        sleep 7
        break
      fi
    done
  fi
}

xfury_tx() {
  # Helper function to broadcast a transaction and supply the necessary args
  # Get module ($1) and specific tx ($2), which forms the tx command
  cmd="$1 $2"
  shift 2

  yes $PASSWORD | xfury tx $cmd \
    --gas-prices $GAS_PRICES \
    --chain-id $CHAIN_ID \
    --broadcast-mode block \
    -y \
    "$@" | jq .
  # The $@ adds any extra arguments to the end
  # --node="$NODE" \
}

xfury_q() {
  xfury q \
    "$@" \
    --output=json | jq .
  # --node="$NODE" \
}

# TRY CATCH implementation following link below
# https://www.xmodulo.com/catch-handle-errors-bash.html
try() {
  [[ $- = *e* ]]
  SAVED_OPT_E=$?
  set +e
}
throw() {
  exit $1
}
catch() {
  export exception_code=$?
  (($SAVED_OPT_E)) && set +e
  return $exception_code
}

full_iid_doc() {
  # Helper function to create a full iid doc => full_iid_doc did address pubkeyBase58
  local DID=$1
  local ADDRESS=$2
  local PUBKEY=$3

  local DID_FULL=$(
    cat <<-END
    {
      "id": "${DID}",
      "controllers": ["${DID}"],
      "verifications": [
        {
          "method": {
            "id": "${DID}",
            "type": "EcdsaSecp256k1VerificationKey2019",
            "controller": "${DID}",
            "publicKeyBase58": "${PUBKEY}"
          },
          "relationships": ["authentication"],
          "context": []
        }
      ],
      "context": [],
      "services": [],
      "accorded_right": [],
      "linked_resources": [],
      "linked_entity": []
    }
END
  )

  echo $DID_FULL
}
