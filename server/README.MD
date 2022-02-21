#### Setup
- Created with: `go mod init server`
- For hot-reloading:
    - Add this to `export PATH=/home/your_name/go/bin:$PATH` to `.bashrc`.
    - Hot reloading docs: `https://github.com/codegangsta/gin`


#### Commands
- Start all commands with `sh run.sh` (i.e. `sh run.sh start`):
    - `start`: Run as is.
    - `hot`: Run with hot-reloading.
    - `build`: Compile.
    - `startb`: Run compiled code.
    - `full`: Build and run compiled code.
- Append the environment name at the end (i.e. `sh run.sh start staging`):
    - If you don't sepecify an environment, `development` is default.
    - Environment names must match file names under `./env`


#### Notes
- Constraints cannot be altered after a table is created. You need to drop the table and recreate the constraint.


#### Todo
Also, rather than run.sh build the commands into the Go executable itself. Check out the Cobra package which is excellent for building CLIs.
