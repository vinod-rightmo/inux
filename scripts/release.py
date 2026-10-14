#!/usr/bin/env python3
"""Build release assets locally or in GitHub Actions; never publish from here."""

import argparse
import hashlib
import os
from pathlib import Path
import re
import shutil
import subprocess
import zipfile

ROOT = Path(__file__).resolve().parents[1]
SKILL = ROOT / "skills" / "inux"
TARGETS = (
    ("darwin", "amd64"),
    ("darwin", "arm64"),
    ("linux", "amd64"),
    ("linux", "arm64"),
    ("windows", "amd64"),
)


def skill_version():
    frontmatter = (SKILL / "SKILL.md").read_text(encoding="utf-8").split("---", 2)[1]
    matches = re.findall(r'^  version: "([^"\n]+)"$', frontmatter, re.MULTILINE)
    if len(matches) != 1 or not re.fullmatch(r"(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)", matches[0]):
        raise ValueError('metadata.version must be a quoted version such as "0.0.1"')
    return matches[0]


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--version", action="store_true", help="print the skill/release version only")
    parser.add_argument("--tag", help="require a tag matching the skill version, for example v0.0.1")
    args = parser.parse_args()
    try:
        version = skill_version()
    except (ValueError, IndexError) as error:
        parser.error(str(error))
    if args.tag is not None and args.tag != "v" + version:
        parser.error(f"tag {args.tag!r} does not match metadata.version {version!r}")
    if args.version:
        print(version)
        return

    output = ROOT / "dist" / "release"
    output.mkdir(parents=True, exist_ok=True)
    assets = []
    for system, architecture in TARGETS:
        suffix = ".exe" if system == "windows" else ""
        binary = output / f"ix-{system}-{architecture}{suffix}"
        environment = dict(os.environ, CGO_ENABLED="0", GOOS=system, GOARCH=architecture)
        subprocess.run(
            ["go", "build", "-mod=readonly", "-trimpath", "-ldflags",
             f"-s -w -X ix/internal.Version={version}", "-o", str(binary), "."],
            cwd=ROOT / "tools" / "ix", env=environment, check=True,
        )
        assets.append(binary)

    package = output / "inux-skill.zip"
    with zipfile.ZipFile(package, "w", zipfile.ZIP_DEFLATED) as archive:
        for path in sorted(SKILL.rglob("*")):
            relative = path.relative_to(SKILL)
            if path.is_file() and path.name != ".DS_Store" and "bin" not in relative.parts:
                archive.write(path, "inux/" + relative.as_posix())
    assets.append(package)
    notices = output / "THIRD_PARTY_NOTICES.txt"
    shutil.copyfile(ROOT / "THIRD_PARTY_NOTICES.txt", notices)
    assets.append(notices)
    checksums = output / "checksums.txt"
    checksums.write_text("".join(
        hashlib.sha256(path.read_bytes()).hexdigest() + "  " + path.name + "\n"
        for path in sorted(assets)
    ), encoding="utf-8")
    print(f"Built Inux {version}: {output}")


if __name__ == "__main__":
    main()
