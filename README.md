Hi there, I made this because I was tired of having multiple versions of my resume on my device.

This is a version control CLI for resumes/curriculum vitae's (CVs) and it wraps git in a way so that you can version control binary files like a pdf, which doesnt happen on git conventionally.
And yes, you can call this a git wrapper because I use 'go-git' under the hood to power the cli functions. 

The commands available with C.V.V.C. as of now are: 
1. cvvc init -> initialises a tempelate resume which you can work on. (Using jake's resume tempelate from Overleaf)
2. cvvc edit -> Opens a local editor at localhost:9090, you can edit your resume here and add/remove/reorder sections and section data as well
3. cvvc status -> Tells you the current branch you are on and the status of commits
4. cvvc commit -m "Message" -> Commits the current state of the resume with a message. Stores in as permenant history
5. cvvc list -> Gives a list of all the past commits you have mad on that resume
6. cvvc branch create BranchName -> create a new branch from your existing one.
7. cvvc switch BranchName -> switched between branches
8. cvvc export -> exports the current resume into a pdf form
9. cvvc diff -> Shows the diff of all the edits made do the resume, compared to the previous commit

What's next to come:
1. Rebase -> Moves your current to a previous commit
2. Selective diff-ing -> Compare between any two versions of the resume, of your choice
3. Tempelates -> currently supports only Jake's resume, would expand to include multiple tempelates that can be selected at time of cvvc init
4. Transform -> convert tempelate of resume from one form to another

If you wish to contribute, PRs are more than welcome. Also if you find any issue feel free to put it up. You can also contact me personally on my email or LinkedIn.    

Currently this is still under development, I will release the 1.0 version sometime in March 2026.
If this brings value to you, do give a star and follow me. I make software that is directly targeted to benefit the end user.

Made with a lot of <3 by Dewashish
