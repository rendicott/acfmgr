PACKAGE DOCUMENTATION

package acfmgr
    import "./acfmgr"

    Package acfmgr is a package to manage entries in an AWS credentials file

    Sample AWS creds file format:

	[default]
	output = json
	region = us-east-1
	aws_access_key_id = QOWIASOVNALKNVCIE
	aws_secret_access_key = zgylMqe64havoaoinweofnviUHqQKYHMGzFMA8CI
	aws_session_token = FQoDYXdzEGYaDNYfEnCsHW/8rG3zpiKwAfS8T...

	[dev-default]
	output = json
	region = us-west-1
	aws_access_key_id = QOWIAADFEGKNVCIE
	aws_secret_access_key = zgylMqaoivnawoeenweofnviUHqQKYHMGzFMA8CI
	aws_session_token = FQoDYXdzEGYaDNYfEnCsanv;oaiwe\iKwAfS8T...

    Adding and removing entries manually is a pain so this package was
    created to assist in programattically adding them once you have sessions
    built from the Golang AWS SDK.

    Calling AssertEntries will delete all entries of that name and only
    rewrite the given entry with the given contents.

    Calling DeleteEntries will delete all entries of that name.

    Sample

	c, err := acfmgr.NewCredFileSession("~/.aws/credentials")
	check(err)
	c.NewEntry("[dev-account-1]", []string{"output = json", "region = us-east-1", "...", ""})
	c.NewEntry("[dev-account-2]", []string{"output = json", "region = us-west-1", "...", ""})
	err = c.AssertEntries()

    Yields:

	[dev-account-1]
	output = json
	region = us-east-1
	...

	[dev-account-2]
	output = json
	region = us-west-1
	...

    While:

	c, err := acfmgr.NewCredFileSession("~/.aws/credentials")
	check(err)
	c.NewEntry("[dev-account-2]", []string{"output = json", "region = us-west-1", "...", ""})
	err = c.DeleteEntries()

    Yields:

	[dev-account-2]
	output = json
	region = us-west-1
	...

TYPES

type CredFile struct {
    // contains filtered or unexported fields
}
    CredFile should be built with the exported NewCredFileSession function.

func NewCredFileSession(filename string) (*CredFile, error)
    NewCredFileSession creates a new interactive credentials file session.
    Needs target filename and returns CredFile obj and err.

func (c *CredFile) AssertEntries() (err error)
    AssertEntries loops through all of the credEntry objs attached to
    CredFile obj and makes sure there is an occurrence with the
    credEntry.name and contents. Existing entries of the same name with
    different contents will be clobbered.

func (c *CredFile) DeleteEntries() (err error)
    DeleteEntries loops through all of the credEntry objs attached to
    CredFile obj and makes sure entries with the same credEntry.name are
    removed. Will remove ALL entries with the same name.

func (c *CredFile) NewEntry(entryName string, entryContents []string)
    NewEntry adds a new credentials entry to the queue to be written or
    deleted with the AssertEntries or DeleteEntries method.

