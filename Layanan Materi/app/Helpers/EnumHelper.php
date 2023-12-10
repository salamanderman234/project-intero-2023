<?php

namespace App\Helpers;
use Illuminate\Support\Facades\DB;

class EnumHelper {
    /**
     * Get enum values from various columns in a given table.
     */
    public static function getEnumValues($table, $column) {
        $enumValues = DB::select("SHOW COLUMNS FROM $table WHERE Field = '{$column}'")[0]->Type;
        preg_match_all("/'([^']+)'/", $enumValues, $matches);
        $enumOptions = $matches[1];
        return $enumOptions;
    }
}