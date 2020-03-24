///////////////////////////////////////////////////////////////////////////////////////
///// IF YOU FOLLOW THESE INSTRUCTIONS ALL THE TESTS SHOULD PASS THE FIRST TIME ///////
///////////////////////////////////////////////////////////////////////////////////////

Instructions on how to run the `dredd` tests

(Prerequisite) Make sure that you have the latest versions of `node` and `npm` and `npx` installed. Then install `dredd` globally using the following command folder:

`npm install dredd --global`

(Running tests) Run the `run_all_tests.sh` script from the `rest_test` directory. This builds the `test.go` 
file, creates the genesis state for the blockchain, starts the blockchain, starts the rest 
server, sends the required transactions to the blockchain, runs all the `dredd` tests, shuts
down the blockchain, cleans up, and propagates up an error code if the tests do not all pass.
**IMPORTANT** - It takes about 3 minutes for the script to run and complete:
**When you run the script 62 tests should pass and zero should fail** 

`./run_all_tests.sh`

(See `test-results` in the `rest_test` for what the `dredd` output should look like.)

(Shutdown) If the script fails or you stop it midway using `ctrl-c` then you should manually stop the blockchain and rest server using the following script. If you let the `run_all_tests.sh` complete
it will automatically shut down the blockchain, rest server and clean up.

`./stopchain.sh`

**DEBUGGING NOTE**: If you start getting `Validator set is different` errors then you need to try starting the chain from scratch (do NOT just use `unsafe-reset-all`, instead use `./stopchain.sh` and then `./run_all_tests.sh`)

