# Instructor Guide: Bring-Your-Own-App hands-on lab

<https://play.instruqt.com/replicated/tracks/support-bundle-own-app-kurl>

## Introduction

This hands-on lab is designed to be delivered by an instructor.  It can be run self-paced, but it benefits greatly from the presence of a guide to discuss and answer questions. The lab is expected to be delivered remotely, with the students sharing their workstation screen with the instructor.

## Instructor's Prerequisites

- A class of not more than 6-8 people.  More than this starts to hinder the discussion-based nature of the lab
- An [Instruqt platform invitation link](#making-an-instruqt-invitation-link)
- A Google Forms link to the post-engagement survey
- A Google Calendar invite with
  - a Zoom link for the session and
  - the Instruqt invitation link
  - the Google Forms link to the survey
- This guide available in a tab

## Setup

### Scheduling the session

Coordinate with the Account Executive who handles the customer account to schedule the session.  Account for approximately 3-4 hours for the session, including time for introductions, breaks, and discussion.

### 1-2 weeks before the session

In the time before the session, the instructor needs to make sure that an Instruqt invitation link has been configured and a Google Calendar invite goes out to the attendees.  Google Calendar should be able to add a Zoom invitation link automatically if you have the extension installed.

#### Google Forms survey

A form already lives on [Google Drive](https://docs.google.com/forms/d/1xZokluh_P1EfLdGyHUSMh7Ngb2bRMB7_4G-XSluETuw/edit?pli=1).  Click "Send" in the upper right corner and get a shareable link to the form, and add it to the Google Calendar invite.

#### Making an Instruqt invitation link

Navigate to the [Invitation Management](https://play.instruqt.com/replicated/invites) screen in the Instruqt platform.  Create a new invitation link with the following settings:

- **Invite Type**: Live Event
- **Public Title**: $Customer / Replicated: BYO-App Labs
- **Public Description**: This is a hands-on lab for $Customer to learn how to troubleshoot Kubernetes problems as they present in your application.  The lab is designed to be delivered by an instructor, and is expected to take 3-4 hours.
- **Select Tracks**: Troubleshoot Your Own App with Replicated - kURL Edition
- **Content Restrictions**: set to expire 1 month after accessing the invite
- **Invitee Restrictions**
  - **Starts On**: set to the date/time of the session
  - **How Many Unique Users**: depends on how the session is coordinated
- **Access Settings**: Anyone leaving their details (recommended)

Click "Create Invite" to create the invitation link.  Copy the link and add it into the calendar invite.

#### (Optional) Configure Hot Start on Instruqt

Hot Start sets up the VMs ahead of time so that attendees don't have to wait as long.  Currently, starting up a session takes about 5-8 minutes from clicking "Start" until you have a usable machine, and we have to install kURL afterward.

Back at the [Invitation Management](https://play.instruqt.com/replicated/invites) screen click on the 3-dot menu next to the invitation link and click "Create Hot Start."  At the next screen, configure the following settings:

- **Select Tracks**: Troubleshoot Your Own App with Replicated - kURL Edition (should already be set if you're coming from the previous step)
- **Number of hot sandboxes per track**: Set to approximately 1.2x the number of attendees, just in case someone needs to start all over again
- **When should Hot Start be available?**: Set to 30m - 1 hour before the session

Click "Create Hot Start"

### Immediately prior to the session

## Running the session

### Introductions

#### Goals

#### Agenda

#### Expectations

### Lab 0: Installing the application

### Lab 1: Pods are not scaled up

### Lab 2: Resource requirements are not met

### Lab 3: Services are not programmed correctly

### Lab 4: Disk full, pods evicted

### Lab 5: DNS failures

### Lab 6: Expired cluster certs (WIP)

### Lab 7: Rook-Ceph failures (WIP)

## Wrap-up

### Survey

### Next steps
