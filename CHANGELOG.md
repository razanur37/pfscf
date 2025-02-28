# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

### Changed

### Removed

### Fixed

## v0.16.5 - 2023-02-11

### Added
- Templates: Added pfs3-01,02,03

## v0.16.4 - 2021-04-03

### Added
- Templates: Added pfs2-12,13,14
- Templates: Added Age of Ashes adventure path

## v0.16.3 - 2021-02-17

### Added
- Templates: Added pfs2.s2-11

### Changed
- Updated external dependencies
- Minor improvements of help texts

## v0.16.2 - 2021-01-17

### Changed
- Templates now distinguish between layout parents and display hierarchie parents. Finally a properly sorted list without having two separate entries for PFS2 season 2 with old and new layout
- Command line arguments provided within the CSV file now automatically remove enclosing quotation marks
- Combine templates for the different tiers of `pfs2.s2-00` into a single template, as the content is identical anyways

### Fixed
- If a command line parameter was set in a CSV file, then it was not possible to overwrite this from the command line in the following `batch fill` call 

## v0.16.1 - 2021-01-11

### Added
- Various PFS2e templates

## v0.16.0 - 2021-01-03

### Added
- Batch mode: Allow to set values for command line parameters from within CSV file
- Batch mode: Allow to set pattern for generated chronicle output filenames to include things like, e.g. character name or society ID.
- PFS2: Scenarios 2-07,2-08; Quests 10-13

### Changed
- Batch mode: Convert most mandatory arguments so that they can be stored within the CSV file now.

## v0.15.2 - 2020-11-22

### Added
- Added new layout for recent scenarios in PFS 2e season 02

### Changed
- The `eventinfo` section will now appear at the top in CSV files generated for PFS 2e.
- Renamed `pfs2.s2` to `pfs2.s2.oldLayout`, as the scenarios switched layout beginning with `pfs2.s05`.

## v0.15.1 - 2020-11-11

### Fixed
- Release artifacts contained relative paths beginning with `.../`. Only the `CHANGELOG.md` and `LICENSE` files were affected, but this still caused errors during unpacking.
- Input from command line and csv file that is UTF8-encoded should now be correctly displayed. So officially now UTF8 and ISO 8859-1 / CP1252 are now supported as input encodings. However, only those UTF8 values are supported that can be displayed by the selected PDF font. So most likely no poo emoji on chronicle sheets.

## v0.15.0 - 2020-11-04

### Changed
- For ease of (template creation) simplicity, items are no longer indexed directly, but their lines are indexed instead. So instead of providing parameter `strikeout_item=1,3`, this might now be something like `strikeout_item_lines=1,3,4` if the third item takes up two lines on the chronicle. Really, really simplifies things for me!
- Finished all PFS2 templates for season 01, quests, season 02 up to s2-04 (incl)

## v0.14.2 - 2020-10-26

### Added
- Print warnings if provided input CSV file does not contain any entries
- Print warnings if provided file names do not have expected file extension
- Templates for PFS2 bounties

### Fixed
- Template files whose filename begins with a dot are now ignored as intended

## v0.14.1 - 2020-10-12

### Added
- Completed some additional scenarios for PFS2 Season 01 that I was asked for.

## v0.14.0 - 2020-10-11

### Added
- New content type `line` to draw lines. Can be used to configure item strikeouts in chronicle templates.
- New content type `strikeout` to draw crosses to the chronicle. These can either be done by providing the center coordinates of the cross in percent and a size in points for things like checkboxes, or by providing two sets of coordinates to span an area for striking out other things like, e.g. a boon.
- Generated CSV files now contain the example values for each parameter in the last column
- First templates for PFS2 Season 02 using the old season 01 layout

### Changed
- Content type `rectangle` now supports an additional field `style`. Currently this accepts two values, `filled` and `strikeout`. With `filled` you get the previous behavior where the rectangle area is filled with the color. With `strikeout` a cross is placed on the area, e.g. to strike out boons. Default is `filled` if no value for field `style` was provided.
- All content types that can directly generate output (i.e. everything except `choice` and `trigger`) will no longer report validation errors in case they have coordinates and sizes that imply no output should be generated. For example a `rectangle` with a width or height of zero will generate no output. Same if e.g. the `fontsize` for a `textentry` is zero. This can be used to have "inactive" content entries in a parent template that will only become active if viable values are provided in a template that inherits from that one, e.g. by providing appropriate preset entries.

## v0.12.0 - 2020-10-05

### Added
- Templates can now have flags. Added `hidden` flag to be able to add technical layout layers in the inheritance hierarchie that should not be displayed with the `template list` command. Will make things easier to handle when different layouts are in place for different seasons (else things would get messy starting with season 3 (yes, 3, not 2))
- New parameter type and content type `multiline`. Currently the content is not automatically split at line end, so this has to be done manually for the moment.

### Changed
- Parameter and content entries of type `choice` can now handle multiple selections. Arguments need to be provided as comma-separated list, e.g. `remove-boons=1,3`. Consequently, the field with name `choice` inside such entries was renamed to `choices`.
- Added layout template `pfs2.sheet_layout1` in hierarchie between templates `pfs2` and `pfs2.s1`, `pfs2.quests` and `pfs2.specials`. This new template contains the description of the "old" layout that was used in PFS2 season 01. For season 02 and others there will be a new layout template `pfs2.sheet_layout2`.
- Output of `template describe` is now grouped and sorted in a similar way to what was already done for the generated CSV files
- Canvases (and their subcanvases) are now considered being "inactive" if their width or height is 0, i.e. if their `x` and `x2` coordinates or `y` and `y2` coordinates are equal. This can be used to define inactive canvases with well-defined names in a parent template, e.g. for striking out boons, that can then easily be filled with the correct coordinates in other templates that inherit from the parent template.

### Fixed
- When calling `pfscf template describe -v <template>`, then the example value for param type `choice` was missing.

## v0.11.3 - 2020-09-29

### Added
- New command line action `open` to have an easy way to open files (e.g. PDF or CSV) with their default application

### Fixed
- Generated CSV files had some typos in comments

## v0.11.2 - 2020-09-27

### Added
- New parameter type 'choice'. Mostly to be used in combination with content type 'choice'.
- New command line flags `--offset-x <value>` and `--offset-y <value>` for the `fill` and `batch fill` commands to fine-tune results in case the result is slightly off.

### Changed
- Template validation now happens in deterministic order, with parent templates being evaluated first
- Renamed content type `textCell` to `text` for simplicity
- Renamed cmd line flag `--exampleValues` to `--examples`
- Renamed cmd line flag `--cellBorder` to `--cell-border`

### Fixed
- Cyclic parent relations in `canvas` entries are now recognized and lead to error messages

## v0.11.1 - 2020-09-25

### Added
- Generated CSV files now include an alphabetically sorted list of all parameters at the end of the file, including the description and example text.
- `pfscf batch fill` now automatically creates the specified output dir if it does not already exist.

### Changed
- Generated CSV file now containes an additional column labelled "GM"

## v0.11.0 - 2020-09-22

### Added
- CSV files generated by the batch command are now auto-opened with the program registered in the OS for CSV files. This behavior can be suppressed with a new command line switch, but is active by default.
- New parameters for pfs2 template for everything in the right chronicle bar.

### Changed
- Parameter configuration in the template files has changed:
  - An additional layer has been introduced to put each parameter entry into a group. For example, parameters `player` and `societyid` are now part of group "Player Information", while parameter `gold` is part of group "Rewards".
  - Parameter entries are o longer sorted alphabetically, but by order of appearance within the CSV file.
  - Together, this should greatly enhance the user experience when filling out a CSV file generated by the pfscf batch mode.

### Fixed
- Coordinates for "sfs" template were broken

## v0.10.1 - 2020-09-21

### Fixed
- Fixed issue with `canvas` entries in `textCell` objects

## v0.10.0 - 2020-09-20

### Added
- Content type `rectangle` now has a new field `transparency` to allow semi-transparent regions. Accepted values are from 0.0 (0%, fully opaque) to 1.0 (100%, fully transparent).
- Documentation is now automatically generated out of the MD files using mkdocs and published to https://razanur37.github.io/pfscf
- The `pfscf fill` command will now automatically open the resulting filled chronicle PDF in the systems default PDF viewer. This can be suppressed with the new `--no-auto-open` command line switch.

### Changed
- The canvas concept. Now totally different than before.
- The option to draw a coordinate grid across the entire page was replaced. Now it is possible to draw a coordinate grid on the canvas with the specified name.

### Removed

### Fixed

## v0.9.2 - 2020-09-13

### Added
- Release now includes a batch file that, when executed via clicking from windows explorer, will open a cmd prompt in the current directory. This should make it easier for windows users to use a command line application like pfscf.
- New subcommand `template search` to search for templates based on search terms.
- New content type `choice` for things like choosing the subtier
- New and improved documentation on usage

### Changed
- Releases for macOS now have "macOS" in their name instead of "Darwin" for clarification
 
## v0.9.1 - 2020-09-11

### Added
- Basic stubs for all PFS2 season 1 scenarios plus quests
- Templates now try to do some basic auto-guessing on possible page margins to reduce the number of cases where values on the produced sheet are misaligned.
- Text is now also automatically shrunk if a textcell is not high enough

### Changed
- Template listing (`pfscf template list`) now shows inheritance relations

## v0.9.0 - 2020-09-10

### Added
- Autoshrink for `textCell`: Automatically reduce font size if text is too wide
- If y2 coordinate is missing or 0, then the cell height is automatically determined via font size
- New template section `parameters`: No longer integrated into content entries. This allows for reusing parameters in a sheet, e.g. GM initials on a Starfinder chronicle
- Content type `rectangle`: Finally drawing colored boxes!
- Content type `trigger`: Can be used to conditionally print other content entries when a specific argument is provided
- Content type `canvas`: Allows to reduce the drawing canvas. Required for easily adapting PFS2 chronicle templates if the right sidebar has different coordinates again
- Chronicle template for Starfinder
- Several chronicle templates for Pathfinder 2 where the right sidebar had a different position in the released sheets

### Changed
- Switched template measurement unit from points to percent
- Empty/missing coordinates are now treated as 0
- Update golang dependencies
- Add missing error check during content generation
- Content entries do no longer have an ID
- Content type `textCell` now also supports static values not related to any passed arguments
- Basically replaced the complete internal data structures and data handling

### Removed
- Content type `societyid`

### Fixed
- Fixed wrong handling of command line arguments in batch mode
- Yaml files with empty sections like `presets:` left the underlying data structure uninitialized due to unexpected behavior in go-yaml module. Workaround was implemented.

## v0.8.0 - 2020-07-30

### Changed
- Heavily restructured internal coding regarding content handling

## v0.7.1 - 2020-07-18

### Changed
- Changed source code structure on disk and goreleaser config

## v0.7.0 - 2020-07-16

### Added
- Documentation on how to write templates
- Batch mode for filling out chronicles

### Changed
- Renamed `x1` and `y1` to `x` and `y`.
- Renamed content `societyid` to `societyId`
- Renamed content `code` to `eventcode`
- Renamed cmd line option `--dummyValues` to `--exampleValues`

### Fixed
- #33: Now correctly only reads files with .yml extension, not with .yml~ or .ymla
- Action `template describe -v` now works again

## v0.6.0 - 2020-07-08

### Added
- New content type `societyid`. This is specifically meant for printing a PFS society id following the pattern `<player_id>-<char_id>`, e.g. 123456-789. This is easier to use than providing both values separately, and also allows better formatting / placement.

### Changed
- Template `pfs2` now provides a `societyid` entry instead of separate `playerid` and `charid` entries. These were removed.

## v0.5.0 - 2020-07-07

### Added
- Template inheritance mechanism
- Mechanism for preset values

### Changed
- Template `pfs2` now uses presets instead of defaults
- Improved error texts

### Removed
- The `default` section is no longer supported / usable

## v0.4.0 - 2020-07-02

### Added
- Align wording: An  `ID` is no longer sometimes called `Name`
- Allow to fill out chronicles with dummy/example values

### Changed
- pfsct is now called pfscf
- Updated pdfcpu from v0.3.3 to v0.3.4
- Use global temp dir now for storing intermediate files
- The `default` section in yaml files was replaced with the section `default` in `presets`

### Fixed
- Now printing filename if an error occurs during reading a yaml file

## v0.3.2 - 2020-06-27

### Added
- Short aliases for cmd line commands, e.g. `f` for `fill`, `t` for `template`
- `verbose` flag and output for `template list` and `template describe`
- Provide example values in verbose output of `template describe`

## v0.3.1 - 2020-06-26

### Added
- First version of `template describe` command

## v0.3.0 - 2020-06-26

### Added
- Proper handling of template files
- First version of `template list` command
- Stubs for other `template` commands

### Fixed
- Allow to execute `pfsct` command when in a different directory

## v0.2.0 - 2020-06-24

### Added
- Check whether required template fields are present
- Mechanism for default values in templates

### Changed
- Configs are now named templates, and thus the `config` subdir was renamed to `template` as well
- Yaml unmarshalling now set to strict

## v0.1 - 2020-06-20

### Added
- First more or less working version that can fill out chronicles for PFS2
