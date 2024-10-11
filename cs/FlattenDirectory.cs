using System;
using System.IO;

class FlattenDirectory
{
    static void FlattenDirectoryStructure(string sourceDir, string targetDir)
    {
        if (!Directory.Exists(sourceDir))
        {
            Console.Error.WriteLine("Error: Source directory does not exist.");
            Environment.Exit(1);
        }

        if (!Directory.Exists(targetDir))
        {
            Directory.CreateDirectory(targetDir);
        }

        foreach (var filePath in Directory.EnumerateFiles(sourceDir, "*", SearchOption.AllDirectories))
        {
            var fileName = Path.GetFileName(filePath);
            var targetPath = Path.Combine(targetDir, fileName);
            File.Copy(filePath, targetPath, true);
        }
        Console.WriteLine($"All files have been extracted to {targetDir} without preserving folder structure.");
    }

    static void Main(string[] args)
    {
        if (args.Length != 2)
        {
            Console.WriteLine("Usage: dotnet FlattenDirectory.dll <source_directory> <target_directory>");
            Environment.Exit(1);
        }
        FlattenDirectoryStructure(args[0], args[1]);
    }
}
