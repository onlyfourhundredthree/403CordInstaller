/*
 * SPDX-License-Identifier: GPL-3.0
 * 403Cord Installer, a cross platform gui/cli app for installing 403Cord
 * Copyright (c) 2023 onlyfourhundredthree and 403Cord contributors
 */

package main

import (
	"image/color"
	"403cordinstaller/buildinfo"
)

const ReleaseUrl = "https://api.github.com/repos/onlyfourhundredthree/403Cord/releases/latest"
const ReleaseUrlFallback = "https://github.com/onlyfourhundredthree/403Cord/releases/latest/download/403cord.zip"
const InstallerReleaseUrl = "https://api.github.com/repos/onlyfourhundredthree/403Cord-Installer/releases/latest"
const InstallerReleaseUrlFallback = "https://github.com/onlyfourhundredthree/403Cord-Installer/releases/latest"

var UserAgent = "403CordInstaller/" + buildinfo.InstallerGitHash + " (https://github.com/onlyfourhundredthree/403Cord-Installer)"

var (
	DiscordGreen  = color.RGBA{R: 0x2D, G: 0x7C, B: 0x46, A: 0xFF}
	DiscordRed    = color.RGBA{R: 0xEC, G: 0x41, B: 0x44, A: 0xFF}
	DiscordBlue   = color.RGBA{R: 0x58, G: 0x65, B: 0xF2, A: 0xFF}
	DiscordYellow = color.RGBA{R: 0xfe, G: 0xe7, B: 0x5c, A: 0xff}
)

var LinuxDiscordNames = []string{
	"Discord",
	"DiscordPTB",
	"DiscordCanary",
	"DiscordDevelopment",
	"discord",
	"discordptb",
	"discordcanary",
	"discorddevelopment",
	"discord-ptb",
	"discord-canary",
	"discord-development",
	"com.discordapp.Discord",
	"com.discordapp.DiscordPTB",
	"com.discordapp.DiscordCanary",
	"com.discordapp.DiscordDevelopment",
}
