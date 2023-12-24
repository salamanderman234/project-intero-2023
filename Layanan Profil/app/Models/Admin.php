<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Admin extends Model
{
    use HasFactory;
    protected $hidden = ["created_at","updated_at","user_id","deleted_at"];
    protected $guarded = ["id", "user_id"];
    public function user()
    {
        return $this->belongsTo(User::class);
    }
}
