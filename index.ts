#!/usr/bin/env node
import execa from "execa";
import os from "os";
import path from "path";
import { name } from "./package.json";

enum NodeArch {
  ARCH_IA32="ia32",
  ARCH_X64="x64",
  ARCH_ARM="arm",
  ARCH_ARM64="arm64"
}

enum GoArch {
  ARCH_386="386",
  ARCH_AMD64="amd64",
  ARCH_ARM="arm",
  ARCH_ARM64="arm64",
}

// Mapping from Node's `process.arch` to Golang's `$GOARCH`
const ARCH_MAPPING: Record<NodeArch, GoArch> = {
  [NodeArch.ARCH_IA32]: GoArch.ARCH_386,
  [NodeArch.ARCH_X64]: GoArch.ARCH_AMD64,
  [NodeArch.ARCH_ARM]: GoArch.ARCH_ARM,
  [NodeArch.ARCH_ARM64]: GoArch.ARCH_ARM64,
};


enum NodePlatform {
  DARWIN="darwin",
  LINUX="linux",
  WIN32="win32",
  FREEBSD="freebsd",
}

enum GoPlatform {
  DARWIN="darwin",
  LINUX="linux",
  WINDOWS="windows",
  FREEBSD="freebsd",
}

// Mapping between Node's `process.platform` to Golang's
const PLATFORM_MAPPING: Record<NodePlatform, GoPlatform> = {
  [NodePlatform.DARWIN]: GoPlatform.DARWIN,
  [NodePlatform.LINUX]: GoPlatform.LINUX,
  [NodePlatform.WIN32]: GoPlatform.WINDOWS,
  [NodePlatform.FREEBSD]: GoPlatform.FREEBSD,
};

function getGoArch(arch: string): GoArch {
  if (arch in ARCH_MAPPING) {
    return ARCH_MAPPING[arch as NodeArch];
  }
  throw new Error("Invalid Arch");
}

function getGoPlatform(platform: string): GoPlatform {
  if (platform in PLATFORM_MAPPING) {
    return PLATFORM_MAPPING[platform as NodePlatform];
  }
  throw new Error("Invalid Platform");
}

function getBinaryPath(goArch: GoArch, goPlatform: GoPlatform): string {
  const folder = `${name}_${goPlatform}_${goArch}`;
  let binary = `${name}`;
  if (goPlatform == GoPlatform.WINDOWS) {
    binary = `${binary}.exe`;
  }
  return path.join(__dirname, "dist", folder, binary);
}

(async function() {

  try {
    console.error("Using NodeJS wrapper...") // TODO pass in as an (hidden argument to Go app)
    const args = process.argv.slice(2)
    const goArch = getGoArch(os.arch());
    const goPlatform = getGoPlatform(os.platform());
    const binaryPath = getBinaryPath(goArch, goPlatform);
    const execution = execa(binaryPath, args);
    execution.stderr?.pipe(process.stderr);
    execution.stdout?.pipe(process.stdout);
    const { exitCode } = await execution; 
    process.exit(exitCode);
  } catch(e: any) {
    // Only print error if there is no stderr from the subprocess
    if (!("stderr" in e)) {
      console.error("NodeJS wrapper failed: ", e);
    }
    process.exit(1);
  }
}());